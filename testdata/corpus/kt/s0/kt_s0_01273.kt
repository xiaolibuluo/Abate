package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10530 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10530 = "registry-17"
    }
}
object Registry10531 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10531 = "registry-24"
    }
}
object Registry10532 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10532 = "registry-31"
    }
}
fun <T, R> List<T>.mapChecked10533(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair10534<A, B>(val first: A, val second: B)

fun combine10534(left: Int, right: Int): Pair10534<Int, Int> {
    return Pair10534(left + 45, right + 43)
}
fun <T> List<T>.secondOrNull10535(): T? = if (size > 1) this[1] else null

fun render10535(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked10536(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
