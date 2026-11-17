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

    public struct Point18770
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 59;
    }

    public static class Grid18770
    {
        public static Point18770 Make(int left, int right)
        {
            return new Point18770 { X = left, Y = right };
        }
    }
    public static int Scale18771<T>(T value) where T : struct
    {
        int number = 66;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public delegate void Notifier18772(int value);

    public sealed class Source18772
    {
        public event Notifier18772 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler18773
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18773(string Label, int Payload);

    public sealed class Handler18773 : IHandler18773
    {
        public void Handle(int value) { }
    }
}
