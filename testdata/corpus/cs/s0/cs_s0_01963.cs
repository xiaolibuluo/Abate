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

    public interface IHandler18290
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18290(string Label, int Payload);

    public sealed class Handler18290 : IHandler18290
    {
        public void Handle(int value) { }
    }
    public struct Point18291
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 24;
    }

    public static class Grid18291
    {
        public static Point18291 Make(int left, int right)
        {
            return new Point18291 { X = left, Y = right };
        }
    }
    public delegate void Notifier18292(int value);

    public sealed class Source18292
    {
        public event Notifier18292 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate18293(IEnumerable<int> source)
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
