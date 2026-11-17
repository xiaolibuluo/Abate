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

    public static int Scale19420<T>(T value) where T : struct
    {
        int number = 31;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 61; }
        return number;
    }
    public struct Point19421
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 38;
    }

    public static class Grid19421
    {
        public static Point19421 Make(int left, int right)
        {
            return new Point19421 { X = left, Y = right };
        }
    }
    public interface IHandler19422
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19422(string Label, int Payload);

    public sealed class Handler19422 : IHandler19422
    {
        public void Handle(int value) { }
    }
    public static int Guard19423(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 52;
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
    public struct Point19424
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 59;
    }

    public static class Grid19424
    {
        public static Point19424 Make(int left, int right)
        {
            return new Point19424 { X = left, Y = right };
        }
    }
    public static int Guard19425(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 66;
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
