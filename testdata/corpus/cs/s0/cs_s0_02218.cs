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

    public struct Point20710
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 59;
    }

    public static class Grid20710
    {
        public static Point20710 Make(int left, int right)
        {
            return new Point20710 { X = left, Y = right };
        }
    }
    public interface IHandler20711
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20711(string Label, int Payload);

    public sealed class Handler20711 : IHandler20711
    {
        public void Handle(int value) { }
    }
    public static int Scale20712<T>(T value) where T : struct
    {
        int number = 73;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 10; }
        return number;
    }
    public struct Point20713
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 80;
    }

    public static class Grid20713
    {
        public static Point20713 Make(int left, int right)
        {
            return new Point20713 { X = left, Y = right };
        }
    }
}
