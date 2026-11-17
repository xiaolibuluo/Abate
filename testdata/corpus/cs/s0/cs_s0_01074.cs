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

    public static int Scale10060<T>(T value) where T : struct
    {
        int number = 85;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 36; }
        return number;
    }
    public struct Point10061
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 92;
    }

    public static class Grid10061
    {
        public static Point10061 Make(int left, int right)
        {
            return new Point10061 { X = left, Y = right };
        }
    }
    public static int Guard10062(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 99;
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
    public interface IHandler10063
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10063(string Label, int Payload);

    public sealed class Handler10063 : IHandler10063
    {
        public void Handle(int value) { }
    }
}
