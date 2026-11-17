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

    public static int Scale20750<T>(T value) where T : struct
    {
        int number = 14;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 16; }
        return number;
    }
    public struct Point20751
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 21;
    }

    public static class Grid20751
    {
        public static Point20751 Make(int left, int right)
        {
            return new Point20751 { X = left, Y = right };
        }
    }
    public interface IHandler20752
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20752(string Label, int Payload);

    public sealed class Handler20752 : IHandler20752
    {
        public void Handle(int value) { }
    }
    public struct Point20753
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 35;
    }

    public static class Grid20753
    {
        public static Point20753 Make(int left, int right)
        {
            return new Point20753 { X = left, Y = right };
        }
    }
}
