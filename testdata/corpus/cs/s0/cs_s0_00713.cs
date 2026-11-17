using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public struct Point6680
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 56;
    }

    public static class Grid6680
    {
        public static Point6680 Make(int left, int right)
        {
            return new Point6680 { X = left, Y = right };
        }
    }
    public delegate void Notifier6681(int value);

    public sealed class Source6681
    {
        public event Notifier6681 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier6682(int value);

    public sealed class Source6682
    {
        public event Notifier6682 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier6683(int value);

    public sealed class Source6683
    {
        public event Notifier6683 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier6684(int value);

    public sealed class Source6684
    {
        public event Notifier6684 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point6685
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 91;
    }

    public static class Grid6685
    {
        public static Point6685 Make(int left, int right)
        {
            return new Point6685 { X = left, Y = right };
        }
    }
}
