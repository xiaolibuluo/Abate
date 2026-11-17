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

    public struct Point16260
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 94;
    }

    public static class Grid16260
    {
        public static Point16260 Make(int left, int right)
        {
            return new Point16260 { X = left, Y = right };
        }
    }
    public delegate void Notifier16261(int value);

    public sealed class Source16261
    {
        public event Notifier16261 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler16262
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16262(string Label, int Payload);

    public sealed class Handler16262 : IHandler16262
    {
        public void Handle(int value) { }
    }
}
