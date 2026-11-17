package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked8430(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair8431<A, B>(val first: A, val second: B)

fun combine8431(left: Int, right: Int): Pair8431<Int, Int> {
    return Pair8431(left + 10, right + 65)
}
object Registry8432 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8432 = "registry-17"
    }
}
fun <T> List<T>.secondOrNull8433(): T? = if (size > 1) this[1] else null

fun render8433(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
