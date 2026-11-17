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

    public interface IHandler12610
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12610(string Label, int Payload);

    public sealed class Handler12610 : IHandler12610
    {
        public void Handle(int value) { }
    }
    public interface IHandler12611
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12611(string Label, int Payload);

    public sealed class Handler12611 : IHandler12611
    {
        public void Handle(int value) { }
    }
    public interface IHandler12612
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12612(string Label, int Payload);

    public sealed class Handler12612 : IHandler12612
    {
        public void Handle(int value) { }
    }
    public interface IHandler12613
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12613(string Label, int Payload);

    public sealed class Handler12613 : IHandler12613
    {
        public void Handle(int value) { }
    }
    public static int Scale12614<T>(T value) where T : struct
    {
        int number = 33;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 35; }
        return number;
    }
    public struct Point12615
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 40;
    }

    public static class Grid12615
    {
        public static Point12615 Make(int left, int right)
        {
            return new Point12615 { X = left, Y = right };
        }
    }
}
