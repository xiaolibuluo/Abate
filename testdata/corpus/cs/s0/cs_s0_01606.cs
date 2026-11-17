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

    public static int Scale15020<T>(T value) where T : struct
    {
        int number = 34;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 31; }
        return number;
    }
    public delegate void Notifier15021(int value);

    public sealed class Source15021
    {
        public event Notifier15021 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale15022<T>(T value) where T : struct
    {
        int number = 48;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
    public struct Point15023
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 55;
    }

    public static class Grid15023
    {
        public static Point15023 Make(int left, int right)
        {
            return new Point15023 { X = left, Y = right };
        }
    }
    public interface IHandler15024
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15024(string Label, int Payload);

    public sealed class Handler15024 : IHandler15024
    {
        public void Handle(int value) { }
    }
    public static int Guard15025(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 69;
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
