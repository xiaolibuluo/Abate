package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17950(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair17951<A, B>(val first: A, val second: B)

fun combine17951(left: Int, right: Int): Pair17951<Int, Int> {
    return Pair17951(left + 67, right + 19)
}
object Registry17952 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17952 = "registry-74"
    }
}
fun <T> List<T>.secondOrNull17953(): T? = if (size > 1) this[1] else null

fun render17953(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step17954")
fun step17954(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 88
}
@JvmName("step17955")
fun step17955(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
object Registry17956 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17956 = "registry-5"
    }
}
fun <T> List<T>.secondOrNull17957(): T? = if (size > 1) this[1] else null

fun render17957(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome17958 {
    data class Success(val payload: String) : Outcome17958()
    data object Failure : Outcome17958()
}

fun describe17958(outcome: Outcome17958): String = when (outcome) {
    is Outcome17958.Success -> outcome.payload
    Outcome17958.Failure -> "failed"
}
