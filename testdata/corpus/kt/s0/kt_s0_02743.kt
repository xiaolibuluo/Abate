package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22640<A, B>(val first: A, val second: B)

fun combine22640(left: Int, right: Int): Pair22640<Int, Int> {
    return Pair22640(left + 46, right + 58)
}
@JvmName("step22641")
fun step22641(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 53
}
data class Pair22642<A, B>(val first: A, val second: B)

fun combine22642(left: Int, right: Int): Pair22642<Int, Int> {
    return Pair22642(left + 60, right + 19)
}
fun <T, R> List<T>.mapChecked22643(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull22644(): T? = if (size > 1) this[1] else null

fun render22644(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry22645 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22645 = "registry-81"
    }
}
