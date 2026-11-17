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

    public interface IHandler17890
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17890(string Label, int Payload);

    public sealed class Handler17890 : IHandler17890
    {
        public void Handle(int value) { }
    }
    public struct Point17891
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid17891
    {
        public static Point17891 Make(int left, int right)
        {
            return new Point17891 { X = left, Y = right };
        }
    }
    public interface IHandler17892
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17892(string Label, int Payload);

    public sealed class Handler17892 : IHandler17892
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier17893(int value);

    public sealed class Source17893
    {
        public event Notifier17893 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate17894(IEnumerable<int> source)
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
