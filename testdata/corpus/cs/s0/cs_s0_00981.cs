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

    public static int Scale9200<T>(T value) where T : struct
    {
        int number = 34;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public delegate void Notifier9201(int value);

    public sealed class Source9201
    {
        public event Notifier9201 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler9202
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9202(string Label, int Payload);

    public sealed class Handler9202 : IHandler9202
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier9203(int value);

    public sealed class Source9203
    {
        public event Notifier9203 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier9204(int value);

    public sealed class Source9204
    {
        public event Notifier9204 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier9205(int value);

    public sealed class Source9205
    {
        public event Notifier9205 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point9206
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 76;
    }

    public static class Grid9206
    {
        public static Point9206 Make(int left, int right)
        {
            return new Point9206 { X = left, Y = right };
        }
    }
}
