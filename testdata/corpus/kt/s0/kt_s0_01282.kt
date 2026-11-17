package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10600<A, B>(val first: A, val second: B)

fun combine10600(left: Int, right: Int): Pair10600<Int, Int> {
    return Pair10600(left + 11, right + 48)
}
object Registry10601 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10601 = "registry-18"
    }
}
fun <T, R> List<T>.mapChecked10602(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull10603(): T? = if (size > 1) this[1] else null

fun render10603(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair10604<A, B>(val first: A, val second: B)

fun combine10604(left: Int, right: Int): Pair10604<Int, Int> {
    return Pair10604(left + 39, right + 31)
}
fun <T> List<T>.secondOrNull10605(): T? = if (size > 1) this[1] else null

fun render10605(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry10606 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10606 = "registry-53"
    }
}
object Registry10607 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10607 = "registry-60"
    }
}
