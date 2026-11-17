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

    public interface IHandler11600
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11600(string Label, int Payload);

    public sealed class Handler11600 : IHandler11600
    {
        public void Handle(int value) { }
    }
    public static int Scale11601<T>(T value) where T : struct
    {
        int number = 57;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 27; }
        return number;
    }
    public struct Point11602
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 64;
    }

    public static class Grid11602
    {
        public static Point11602 Make(int left, int right)
        {
            return new Point11602 { X = left, Y = right };
        }
    }
    public static int Scale11603<T>(T value) where T : struct
    {
        int number = 71;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public static int Scale11604<T>(T value) where T : struct
    {
        int number = 78;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 60; }
        return number;
    }
    public static int Guard11605(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 85;
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
}
