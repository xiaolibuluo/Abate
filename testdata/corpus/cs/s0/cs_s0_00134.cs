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

    public struct Point1290
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 33;
    }

    public static class Grid1290
    {
        public static Point1290 Make(int left, int right)
        {
            return new Point1290 { X = left, Y = right };
        }
    }
    public delegate void Notifier1291(int value);

    public sealed class Source1291
    {
        public event Notifier1291 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler1292
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope1292(string Label, int Payload);

    public sealed class Handler1292 : IHandler1292
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate1293(IEnumerable<int> source)
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
}
