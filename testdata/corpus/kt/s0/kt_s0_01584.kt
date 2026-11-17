package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step13130")
fun step13130(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
object Registry13131 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13131 = "registry-9"
    }
}
sealed class Outcome13132 {
    data class Success(val payload: String) : Outcome13132()
    data object Failure : Outcome13132()
}

fun describe13132(outcome: Outcome13132): String = when (outcome) {
    is Outcome13132.Success -> outcome.payload
    Outcome13132.Failure -> "failed"
}
fun <T> List<T>.secondOrNull13133(): T? = if (size > 1) this[1] else null

fun render13133(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13134 {
    data class Success(val payload: String) : Outcome13134()
    data object Failure : Outcome13134()
}

fun describe13134(outcome: Outcome13134): String = when (outcome) {
    is Outcome13134.Success -> outcome.payload
    Outcome13134.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13135(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry13136 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13136 = "registry-44"
    }
}
sealed class Outcome13137 {
    data class Success(val payload: String) : Outcome13137()
    data object Failure : Outcome13137()
}

fun describe13137(outcome: Outcome13137): String = when (outcome) {
    is Outcome13137.Success -> outcome.payload
    Outcome13137.Failure -> "failed"
}
fun <T> List<T>.secondOrNull13138(): T? = if (size > 1) this[1] else null

fun render13138(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
