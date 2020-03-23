import * as React from 'react'
import { Container } from 'unstated-typescript'

export interface ContainerClass<
    TState extends object,
    TContainer extends Container<TState> = Container<TState>
    > {
    new(...args: any[]): TContainer;
}
export type ContainerType = ContainerClass<any> | Container<any>;

type ContainerInstancesType<TContainers extends [ContainerType, ...ContainerType[]]> = {
    [K in keyof TContainers]: TContainers[K] extends ContainerClass<
        any,
        infer TContainer
    >
    ? TContainer
    : TContainers[K]
};

export interface SubscribeProps<TContainers extends [ContainerType, ...ContainerType[]]> {
    to: TContainers;
    children(
        ...instances: ContainerInstancesType<TContainers>
    ): React.ReactNode;
}

export class Subscribe<TContainers extends [ContainerType, ...ContainerType[]]>
    extends React.Component<SubscribeProps<TContainers>> { }

export interface ProviderProps {
    inject?: [Container<any>, ...Container<any>[]];
    children: React.ReactNode;
}