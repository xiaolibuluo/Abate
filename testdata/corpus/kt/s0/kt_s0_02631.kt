package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21720 {
    data class Success(val payload: String) : Outcome21720()
    data object Failure : Outcome21720()
}

fun describe21720(outcome: Outcome21720): String = when (outcome) {
    is Outcome21720.Success -> outcome.payload
    Outcome21720.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21721(): T? = if (size > 1) this[1] else null

fun render21721(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry21722 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21722 = "registry-28"
    }
}
@JvmName("step21723")
fun step21723(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
fun <T, R> List<T>.mapChecked21724(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull21725(): T? = if (size > 1) this[1] else null

fun render21725(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome21726 {
    data class Success(val payload: String) : Outcome21726()
    data object Failure : Outcome21726()
}

fun describe21726(outcome: Outcome21726): String = when (outcome) {
    is Outcome21726.Success -> outcome.payload
    Outcome21726.Failure -> "failed"
}
