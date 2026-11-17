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

    public struct Point5180
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 46;
    }

    public static class Grid5180
    {
        public static Point5180 Make(int left, int right)
        {
            return new Point5180 { X = left, Y = right };
        }
    }
    public struct Point5181
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 53;
    }

    public static class Grid5181
    {
        public static Point5181 Make(int left, int right)
        {
            return new Point5181 { X = left, Y = right };
        }
    }
    public delegate void Notifier5182(int value);

    public sealed class Source5182
    {
        public event Notifier5182 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler5183
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5183(string Label, int Payload);

    public sealed class Handler5183 : IHandler5183
    {
        public void Handle(int value) { }
    }
}
