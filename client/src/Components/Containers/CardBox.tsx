import {
  styled,
} from '@mui/material';
import {
  Box,
  BoxProps,
} from '@mui/system';
import {
  useMemo,
} from 'react';
import CardSkeleton from '../Skeletons/CardSkeleton';

export const CardBox = styled(Box)<BoxProps>(({ theme }) => ({
  display: 'grid',
  columnGap: '12px',
  rowGap: '12px',
  alignItems: 'start',
  paddingBottom: '128px', // Add some extra space so the FAB doesn't overlay the actions.
  gridAutoRows: 'min-content',
  width: '100%',
  boxSizing: 'border-box',
  [theme.breakpoints.up('md')]: {
    justifyContent: 'center',
    gridTemplateColumns: '320px 320px',
  },
  [theme.breakpoints.down('sm')]: {
    gridTemplateColumns: 'auto',
  }
}));

/**
 * See {@link LoadingCardBox}.
 */
export type LoadingCardBoxProps = BoxProps & {
  /**
   * If true, then render some card skeletons rather than any cards.
   */
  isLoading: boolean;

  /**
   * The number of card skeletons to render when loading. The default value is 2.
   */
  skeletonCount?: number;
}

/**
 * A component used to render card skeletons when {@link isLoading} is true.
 * 
 * @param props See {@link LoadingCardBoxProps}.
 * @returns The component.
 */
export function LoadingCardBox(props: React.PropsWithChildren<LoadingCardBoxProps>) {
  const {
    children,
    isLoading,
    skeletonCount = 2,
  } = props;

  const skeletons = useMemo(() => {
    let skeletonArray: JSX.Element[] = [];
    for (let i = 0; i < skeletonCount; i++) {
      skeletonArray = [...skeletonArray, (<CardSkeleton />)];
    }
    return skeletonArray;
  }, [skeletonCount]);

  const elementsToDisplay = useMemo(() => isLoading
    ? skeletons
    : children, [children, isLoading]
  );

  return (
    <CardBox>
      {elementsToDisplay}
    </CardBox>
  );
}
