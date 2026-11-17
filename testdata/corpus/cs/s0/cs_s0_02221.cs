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

    public struct Point20740
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid20740
    {
        public static Point20740 Make(int left, int right)
        {
            return new Point20740 { X = left, Y = right };
        }
    }
    public static int Guard20741(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 8;
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
    public interface IHandler20742
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20742(string Label, int Payload);

    public sealed class Handler20742 : IHandler20742
    {
        public void Handle(int value) { }
    }
    public static int Scale20743<T>(T value) where T : struct
    {
        int number = 22;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 42; }
        return number;
    }
}
