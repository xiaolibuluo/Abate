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

    public async Task<int> Aggregate23670(IEnumerable<int> source)
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
    public interface IHandler23671
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23671(string Label, int Payload);

    public sealed class Handler23671 : IHandler23671
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier23672(int value);

    public sealed class Source23672
    {
        public event Notifier23672 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point23673
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 48;
    }

    public static class Grid23673
    {
        public static Point23673 Make(int left, int right)
        {
            return new Point23673 { X = left, Y = right };
        }
    }
}
