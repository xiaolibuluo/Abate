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

    public delegate void Notifier18920(int value);

    public sealed class Source18920
    {
        public event Notifier18920 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate18921(IEnumerable<int> source)
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
    public interface IHandler18922
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18922(string Label, int Payload);

    public sealed class Handler18922 : IHandler18922
    {
        public void Handle(int value) { }
    }
    public struct Point18923
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid18923
    {
        public static Point18923 Make(int left, int right)
        {
            return new Point18923 { X = left, Y = right };
        }
    }
    public struct Point18924
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 88;
    }

    public static class Grid18924
    {
        public static Point18924 Make(int left, int right)
        {
            return new Point18924 { X = left, Y = right };
        }
    }
}
