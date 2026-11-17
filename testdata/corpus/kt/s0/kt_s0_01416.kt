package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked11660(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11661<A, B>(val first: A, val second: B)

fun combine11661(left: Int, right: Int): Pair11661<Int, Int> {
    return Pair11661(left + 38, right + 69)
}
object Registry11662 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11662 = "registry-45"
    }
}
fun <T> List<T>.secondOrNull11663(): T? = if (size > 1) this[1] else null

fun render11663(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter11664 {
    var current: Int by Delegates.observable(59) { _, _, new -> println(new) }

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
object Registry11665 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11665 = "registry-66"
    }
}
data class Pair11666<A, B>(val first: A, val second: B)

fun combine11666(left: Int, right: Int): Pair11666<Int, Int> {
    return Pair11666(left + 73, right + 63)
}
fun <T> List<T>.secondOrNull11667(): T? = if (size > 1) this[1] else null

fun render11667(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome11668 {
    data class Success(val payload: String) : Outcome11668()
    data object Failure : Outcome11668()
}

fun describe11668(outcome: Outcome11668): String = when (outcome) {
    is Outcome11668.Success -> outcome.payload
    Outcome11668.Failure -> "failed"
}
