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

    public interface IHandler16250
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16250(string Label, int Payload);

    public sealed class Handler16250 : IHandler16250
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier16251(int value);

    public sealed class Source16251
    {
        public event Notifier16251 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point16252
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 95;
    }

    public static class Grid16252
    {
        public static Point16252 Make(int left, int right)
        {
            return new Point16252 { X = left, Y = right };
        }
    }
}
