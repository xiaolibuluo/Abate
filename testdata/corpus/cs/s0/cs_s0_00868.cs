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

    public delegate void Notifier8140(int value);

    public sealed class Source8140
    {
        public event Notifier8140 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point8141
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 21;
    }

    public static class Grid8141
    {
        public static Point8141 Make(int left, int right)
        {
            return new Point8141 { X = left, Y = right };
        }
    }
    public struct Point8142
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 28;
    }

    public static class Grid8142
    {
        public static Point8142 Make(int left, int right)
        {
            return new Point8142 { X = left, Y = right };
        }
    }
    public struct Point8143
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 35;
    }

    public static class Grid8143
    {
        public static Point8143 Make(int left, int right)
        {
            return new Point8143 { X = left, Y = right };
        }
    }
    public interface IHandler8144
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8144(string Label, int Payload);

    public sealed class Handler8144 : IHandler8144
    {
        public void Handle(int value) { }
    }
}
