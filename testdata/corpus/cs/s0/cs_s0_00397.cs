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

    public interface IHandler3750
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3750(string Label, int Payload);

    public sealed class Handler3750 : IHandler3750
    {
        public void Handle(int value) { }
    }
    public interface IHandler3751
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3751(string Label, int Payload);

    public sealed class Handler3751 : IHandler3751
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier3752(int value);

    public sealed class Source3752
    {
        public event Notifier3752 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point3753
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 51;
    }

    public static class Grid3753
    {
        public static Point3753 Make(int left, int right)
        {
            return new Point3753 { X = left, Y = right };
        }
    }
    public struct Point3754
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 58;
    }

    public static class Grid3754
    {
        public static Point3754 Make(int left, int right)
        {
            return new Point3754 { X = left, Y = right };
        }
    }
}
