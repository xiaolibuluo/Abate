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

    public static int Scale13070<T>(T value) where T : struct
    {
        int number = 21;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 69; }
        return number;
    }
    public interface IHandler13071
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13071(string Label, int Payload);

    public sealed class Handler13071 : IHandler13071
    {
        public void Handle(int value) { }
    }
    public struct Point13072
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 35;
    }

    public static class Grid13072
    {
        public static Point13072 Make(int left, int right)
        {
            return new Point13072 { X = left, Y = right };
        }
    }
    public static int Guard13073(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 42;
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
    public static int Scale13074<T>(T value) where T : struct
    {
        int number = 49;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public delegate void Notifier13075(int value);

    public sealed class Source13075
    {
        public event Notifier13075 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
