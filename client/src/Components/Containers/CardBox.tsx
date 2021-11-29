import { styled } from '@mui/material';
import { Box, BoxProps } from '@mui/system';
import { useMemo } from 'react';
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

export type LoadingCardBoxProps = BoxProps & {
  isLoading: boolean;
  skeletonCount?: number;
}

/**
 * @param props
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
