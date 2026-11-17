package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6700(): T? = if (size > 1) this[1] else null

fun render6700(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull6701(): T? = if (size > 1) this[1] else null

fun render6701(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6702 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6702 = "registry-96"
    }
}
object Registry6703 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6703 = "registry-6"
    }
}
sealed class Outcome6704 {
    data class Success(val payload: String) : Outcome6704()
    data object Failure : Outcome6704()
}

fun describe6704(outcome: Outcome6704): String = when (outcome) {
    is Outcome6704.Success -> outcome.payload
    Outcome6704.Failure -> "failed"
}
sealed class Outcome6705 {
    data class Success(val payload: String) : Outcome6705()
    data object Failure : Outcome6705()
}

fun describe6705(outcome: Outcome6705): String = when (outcome) {
    is Outcome6705.Success -> outcome.payload
    Outcome6705.Failure -> "failed"
}
fun <T> List<T>.secondOrNull6706(): T? = if (size > 1) this[1] else null

fun render6706(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
