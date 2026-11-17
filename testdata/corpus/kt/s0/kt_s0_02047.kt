package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked16930(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry16931 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16931 = "registry-99"
    }
}
fun <T> List<T>.secondOrNull16932(): T? = if (size > 1) this[1] else null

fun render16932(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull16933(): T? = if (size > 1) this[1] else null

fun render16933(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter16934 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
data class Pair16935<A, B>(val first: A, val second: B)

fun combine16935(left: Int, right: Int): Pair16935<Int, Int> {
    return Pair16935(left + 30, right + 20)
}
sealed class Outcome16936 {
    data class Success(val payload: String) : Outcome16936()
    data object Failure : Outcome16936()
}

fun describe16936(outcome: Outcome16936): String = when (outcome) {
    is Outcome16936.Success -> outcome.payload
    Outcome16936.Failure -> "failed"
}
