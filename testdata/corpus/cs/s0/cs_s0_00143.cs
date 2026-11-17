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

    public struct Point1370
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 40;
    }

    public static class Grid1370
    {
        public static Point1370 Make(int left, int right)
        {
            return new Point1370 { X = left, Y = right };
        }
    }
    public struct Point1371
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 47;
    }

    public static class Grid1371
    {
        public static Point1371 Make(int left, int right)
        {
            return new Point1371 { X = left, Y = right };
        }
    }
    public static int Scale1372<T>(T value) where T : struct
    {
        int number = 54;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 14; }
        return number;
    }
    public interface IHandler1373
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope1373(string Label, int Payload);

    public sealed class Handler1373 : IHandler1373
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier1374(int value);

    public sealed class Source1374
    {
        public event Notifier1374 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point1375
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 75;
    }

    public static class Grid1375
    {
        public static Point1375 Make(int left, int right)
        {
            return new Point1375 { X = left, Y = right };
        }
    }
}
