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

    public static int Scale19900<T>(T value) where T : struct
    {
        int number = 73;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 31; }
        return number;
    }
    public delegate void Notifier19901(int value);

    public sealed class Source19901
    {
        public event Notifier19901 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard19902(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 87;
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
    public interface IHandler19903
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19903(string Label, int Payload);

    public sealed class Handler19903 : IHandler19903
    {
        public void Handle(int value) { }
    }
    public struct Point19904
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 101;
    }

    public static class Grid19904
    {
        public static Point19904 Make(int left, int right)
        {
            return new Point19904 { X = left, Y = right };
        }
    }
    public interface IHandler19905
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19905(string Label, int Payload);

    public sealed class Handler19905 : IHandler19905
    {
        public void Handle(int value) { }
    }
}
