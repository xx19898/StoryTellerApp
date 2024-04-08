import type { Meta, StoryObj } from '@storybook/react';
import ResponsiveTextArea from './ResponsiveTextArea';

// More on how to set up stories at: https://storybook.js.org/docs/react/writing-stories/introduction
const meta = {
  title: 'TextArea that shrinks and grows with its contents',
  component: ResponsiveTextArea,
  tags: ['autodocs'],
  argTypes: {
  },
} satisfies Meta<typeof ResponsiveTextArea>;

export default meta;
type Story = StoryObj<typeof ResponsiveTextArea>;

export const Default: Story = {
  args: {}
};
