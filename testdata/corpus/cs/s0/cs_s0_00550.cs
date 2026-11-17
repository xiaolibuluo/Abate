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

    public struct Point5150
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 7;
    }

    public static class Grid5150
    {
        public static Point5150 Make(int left, int right)
        {
            return new Point5150 { X = left, Y = right };
        }
    }
    public interface IHandler5151
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5151(string Label, int Payload);

    public sealed class Handler5151 : IHandler5151
    {
        public void Handle(int value) { }
    }
    public static int Guard5152(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 21;
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
    public delegate void Notifier5153(int value);

    public sealed class Source5153
    {
        public event Notifier5153 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
