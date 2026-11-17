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

    public static int Scale5600<T>(T value) where T : struct
    {
        int number = 10;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 25; }
        return number;
    }
    public interface IHandler5601
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5601(string Label, int Payload);

    public sealed class Handler5601 : IHandler5601
    {
        public void Handle(int value) { }
    }
    public static int Scale5602<T>(T value) where T : struct
    {
        int number = 24;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public struct Point5603
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 31;
    }

    public static class Grid5603
    {
        public static Point5603 Make(int left, int right)
        {
            return new Point5603 { X = left, Y = right };
        }
    }
}
