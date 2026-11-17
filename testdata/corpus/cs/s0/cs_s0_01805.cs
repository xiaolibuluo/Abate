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

    public interface IHandler16840
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16840(string Label, int Payload);

    public sealed class Handler16840 : IHandler16840
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate16841(IEnumerable<int> source)
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
    public delegate void Notifier16842(int value);

    public sealed class Source16842
    {
        public event Notifier16842 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point16843
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 93;
    }

    public static class Grid16843
    {
        public static Point16843 Make(int left, int right)
        {
            return new Point16843 { X = left, Y = right };
        }
    }
}
