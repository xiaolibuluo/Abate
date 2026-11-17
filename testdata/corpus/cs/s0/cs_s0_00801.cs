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

    public interface IHandler7520
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7520(string Label, int Payload);

    public sealed class Handler7520 : IHandler7520
    {
        public void Handle(int value) { }
    }
    public static int Scale7521<T>(T value) where T : struct
    {
        int number = 88;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 38; }
        return number;
    }
    public struct Point7522
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 95;
    }

    public static class Grid7522
    {
        public static Point7522 Make(int left, int right)
        {
            return new Point7522 { X = left, Y = right };
        }
    }
}
