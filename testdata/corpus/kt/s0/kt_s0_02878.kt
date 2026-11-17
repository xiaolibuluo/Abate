package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry23740 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23740 = "registry-21"
    }
}
data class Pair23741<A, B>(val first: A, val second: B)

fun combine23741(left: Int, right: Int): Pair23741<Int, Int> {
    return Pair23741(left + 28, right + 46)
}
fun <T> List<T>.secondOrNull23742(): T? = if (size > 1) this[1] else null

fun render23742(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry23743 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23743 = "registry-42"
    }
}
object Registry23744 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23744 = "registry-49"
    }
}
fun <T, R> List<T>.mapChecked23745(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step23746")
fun step23746(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
object Registry23747 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23747 = "registry-70"
    }
}
data class Pair23748<A, B>(val first: A, val second: B)

fun combine23748(left: Int, right: Int): Pair23748<Int, Int> {
    return Pair23748(left + 77, right + 62)
}
