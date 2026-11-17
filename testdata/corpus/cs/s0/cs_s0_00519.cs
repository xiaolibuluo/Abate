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

    public interface IHandler4860
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4860(string Label, int Payload);

    public sealed class Handler4860 : IHandler4860
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier4861(int value);

    public sealed class Source4861
    {
        public event Notifier4861 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale4862<T>(T value) where T : struct
    {
        int number = 32;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
    public struct Point4863
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 39;
    }

    public static class Grid4863
    {
        public static Point4863 Make(int left, int right)
        {
            return new Point4863 { X = left, Y = right };
        }
    }
    public interface IHandler4864
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4864(string Label, int Payload);

    public sealed class Handler4864 : IHandler4864
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier4865(int value);

    public sealed class Source4865
    {
        public event Notifier4865 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler4866
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4866(string Label, int Payload);

    public sealed class Handler4866 : IHandler4866
    {
        public void Handle(int value) { }
    }
}
