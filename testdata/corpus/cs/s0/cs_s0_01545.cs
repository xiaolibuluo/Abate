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

    public static int Scale14470<T>(T value) where T : struct
    {
        int number = 95;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 12; }
        return number;
    }
    public static int Guard14471(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 5;
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
    public struct Point14472
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 12;
    }

    public static class Grid14472
    {
        public static Point14472 Make(int left, int right)
        {
            return new Point14472 { X = left, Y = right };
        }
    }
    public static int Guard14473(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 19;
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
    public interface IHandler14474
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14474(string Label, int Payload);

    public sealed class Handler14474 : IHandler14474
    {
        public void Handle(int value) { }
    }
}
