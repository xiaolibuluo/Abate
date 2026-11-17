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

    public struct Point11560
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 95;
    }

    public static class Grid11560
    {
        public static Point11560 Make(int left, int right)
        {
            return new Point11560 { X = left, Y = right };
        }
    }
    public interface IHandler11561
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11561(string Label, int Payload);

    public sealed class Handler11561 : IHandler11561
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate11562(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public interface IHandler11563
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11563(string Label, int Payload);

    public sealed class Handler11563 : IHandler11563
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier11564(int value);

    public sealed class Source11564
    {
        public event Notifier11564 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
