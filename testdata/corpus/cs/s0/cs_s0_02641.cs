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

    public struct Point24710
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 21;
    }

    public static class Grid24710
    {
        public static Point24710 Make(int left, int right)
        {
            return new Point24710 { X = left, Y = right };
        }
    }
    public interface IHandler24711
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24711(string Label, int Payload);

    public sealed class Handler24711 : IHandler24711
    {
        public void Handle(int value) { }
    }
    public interface IHandler24712
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24712(string Label, int Payload);

    public sealed class Handler24712 : IHandler24712
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate24713(IEnumerable<int> source)
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
    public delegate void Notifier24714(int value);

    public sealed class Source24714
    {
        public event Notifier24714 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
