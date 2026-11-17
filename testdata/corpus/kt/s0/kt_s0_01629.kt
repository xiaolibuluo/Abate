package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step13460")
fun step13460(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
data class Pair13461<A, B>(val first: A, val second: B)

fun combine13461(left: Int, right: Int): Pair13461<Int, Int> {
    return Pair13461(left + 50, right + 48)
}
fun <T> List<T>.secondOrNull13462(): T? = if (size > 1) this[1] else null

fun render13462(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry13463 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13463 = "registry-64"
    }
}
fun <T> List<T>.secondOrNull13464(): T? = if (size > 1) this[1] else null

fun render13464(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked13465(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry13466 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13466 = "registry-85"
    }
}
