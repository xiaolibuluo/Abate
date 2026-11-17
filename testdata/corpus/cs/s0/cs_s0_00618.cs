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

    public struct Point5770
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 37;
    }

    public static class Grid5770
    {
        public static Point5770 Make(int left, int right)
        {
            return new Point5770 { X = left, Y = right };
        }
    }
    public static int Guard5771(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 44;
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
    public delegate void Notifier5772(int value);

    public sealed class Source5772
    {
        public event Notifier5772 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler5773
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5773(string Label, int Payload);

    public sealed class Handler5773 : IHandler5773
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier5774(int value);

    public sealed class Source5774
    {
        public event Notifier5774 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier5775(int value);

    public sealed class Source5775
    {
        public event Notifier5775 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
