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

    public interface IHandler4810
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4810(string Label, int Payload);

    public sealed class Handler4810 : IHandler4810
    {
        public void Handle(int value) { }
    }
    public struct Point4811
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 57;
    }

    public static class Grid4811
    {
        public static Point4811 Make(int left, int right)
        {
            return new Point4811 { X = left, Y = right };
        }
    }
    public interface IHandler4812
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4812(string Label, int Payload);

    public sealed class Handler4812 : IHandler4812
    {
        public void Handle(int value) { }
    }
    public struct Point4813
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 71;
    }

    public static class Grid4813
    {
        public static Point4813 Make(int left, int right)
        {
            return new Point4813 { X = left, Y = right };
        }
    }
    public static int Scale4814<T>(T value) where T : struct
    {
        int number = 78;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 65; }
        return number;
    }
    public interface IHandler4815
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4815(string Label, int Payload);

    public sealed class Handler4815 : IHandler4815
    {
        public void Handle(int value) { }
    }
}
