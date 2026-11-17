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

    public delegate void Notifier20240(int value);

    public sealed class Source20240
    {
        public event Notifier20240 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point20241
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 37;
    }

    public static class Grid20241
    {
        public static Point20241 Make(int left, int right)
        {
            return new Point20241 { X = left, Y = right };
        }
    }
    public struct Point20242
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 44;
    }

    public static class Grid20242
    {
        public static Point20242 Make(int left, int right)
        {
            return new Point20242 { X = left, Y = right };
        }
    }
    public interface IHandler20243
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20243(string Label, int Payload);

    public sealed class Handler20243 : IHandler20243
    {
        public void Handle(int value) { }
    }
}
