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

    public interface IHandler20360
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20360(string Label, int Payload);

    public sealed class Handler20360 : IHandler20360
    {
        public void Handle(int value) { }
    }
    public static int Guard20361(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 96;
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
    public delegate void Notifier20362(int value);

    public sealed class Source20362
    {
        public event Notifier20362 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier20363(int value);

    public sealed class Source20363
    {
        public event Notifier20363 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point20364
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 20;
    }

    public static class Grid20364
    {
        public static Point20364 Make(int left, int right)
        {
            return new Point20364 { X = left, Y = right };
        }
    }
    public delegate void Notifier20365(int value);

    public sealed class Source20365
    {
        public event Notifier20365 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
