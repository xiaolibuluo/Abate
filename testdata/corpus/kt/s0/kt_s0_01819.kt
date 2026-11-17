package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry14960 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14960 = "registry-53"
    }
}
data class Pair14961<A, B>(val first: A, val second: B)

fun combine14961(left: Int, right: Int): Pair14961<Int, Int> {
    return Pair14961(left + 60, right + 61)
}
fun <T, R> List<T>.mapChecked14962(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step14963")
fun step14963(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
fun <T> List<T>.secondOrNull14964(): T? = if (size > 1) this[1] else null

fun render14964(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14965 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14965 = "registry-88"
    }
}
sealed class Outcome14966 {
    data class Success(val payload: String) : Outcome14966()
    data object Failure : Outcome14966()
}

fun describe14966(outcome: Outcome14966): String = when (outcome) {
    is Outcome14966.Success -> outcome.payload
    Outcome14966.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked14967(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
