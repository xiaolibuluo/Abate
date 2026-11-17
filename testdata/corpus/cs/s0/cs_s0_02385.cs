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

    public static int Scale22270<T>(T value) where T : struct
    {
        int number = 50;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 43; }
        return number;
    }
    public async Task<int> Aggregate22271(IEnumerable<int> source)
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
    public delegate void Notifier22272(int value);

    public sealed class Source22272
    {
        public event Notifier22272 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point22273
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 71;
    }

    public static class Grid22273
    {
        public static Point22273 Make(int left, int right)
        {
            return new Point22273 { X = left, Y = right };
        }
    }
    public interface IHandler22274
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22274(string Label, int Payload);

    public sealed class Handler22274 : IHandler22274
    {
        public void Handle(int value) { }
    }
}
