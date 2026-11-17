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

    public struct Point4730
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 43;
    }

    public static class Grid4730
    {
        public static Point4730 Make(int left, int right)
        {
            return new Point4730 { X = left, Y = right };
        }
    }
    public struct Point4731
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 50;
    }

    public static class Grid4731
    {
        public static Point4731 Make(int left, int right)
        {
            return new Point4731 { X = left, Y = right };
        }
    }
    public static int Scale4732<T>(T value) where T : struct
    {
        int number = 57;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 48; }
        return number;
    }
    public static int Guard4733(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 64;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public interface IHandler4734
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4734(string Label, int Payload);

    public sealed class Handler4734 : IHandler4734
    {
        public void Handle(int value) { }
    }
}
