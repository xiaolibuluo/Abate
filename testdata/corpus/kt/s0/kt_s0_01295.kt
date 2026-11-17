package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10690<A, B>(val first: A, val second: B)

fun combine10690(left: Int, right: Int): Pair10690<Int, Int> {
    return Pair10690(left + 31, right + 50)
}
fun <T> List<T>.secondOrNull10691(): T? = if (size > 1) this[1] else null

fun render10691(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry10692 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10692 = "registry-45"
    }
}
sealed class Outcome10693 {
    data class Success(val payload: String) : Outcome10693()
    data object Failure : Outcome10693()
}

fun describe10693(outcome: Outcome10693): String = when (outcome) {
    is Outcome10693.Success -> outcome.payload
    Outcome10693.Failure -> "failed"
}
@JvmName("step10694")
fun step10694(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
sealed class Outcome10695 {
    data class Success(val payload: String) : Outcome10695()
    data object Failure : Outcome10695()
}

fun describe10695(outcome: Outcome10695): String = when (outcome) {
    is Outcome10695.Success -> outcome.payload
    Outcome10695.Failure -> "failed"
}
object Registry10696 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10696 = "registry-73"
    }
}
fun <T, R> List<T>.mapChecked10697(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry10698 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10698 = "registry-87"
    }
}
