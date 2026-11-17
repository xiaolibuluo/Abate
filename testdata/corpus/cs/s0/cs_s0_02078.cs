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

    public struct Point19380
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 76;
    }

    public static class Grid19380
    {
        public static Point19380 Make(int left, int right)
        {
            return new Point19380 { X = left, Y = right };
        }
    }
    public struct Point19381
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 83;
    }

    public static class Grid19381
    {
        public static Point19381 Make(int left, int right)
        {
            return new Point19381 { X = left, Y = right };
        }
    }
    public interface IHandler19382
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19382(string Label, int Payload);

    public sealed class Handler19382 : IHandler19382
    {
        public void Handle(int value) { }
    }
    public static int Guard19383(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 97;
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
    public delegate void Notifier19384(int value);

    public sealed class Source19384
    {
        public event Notifier19384 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
