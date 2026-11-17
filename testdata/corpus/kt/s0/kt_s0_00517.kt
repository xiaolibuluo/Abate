package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome4180 {
    data class Success(val payload: String) : Outcome4180()
    data object Failure : Outcome4180()
}

fun describe4180(outcome: Outcome4180): String = when (outcome) {
    is Outcome4180.Success -> outcome.payload
    Outcome4180.Failure -> "failed"
}
fun <T> List<T>.secondOrNull4181(): T? = if (size > 1) this[1] else null

fun render4181(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry4182 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4182 = "registry-21"
    }
}
object Registry4183 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4183 = "registry-28"
    }
}
sealed class Outcome4184 {
    data class Success(val payload: String) : Outcome4184()
    data object Failure : Outcome4184()
}

fun describe4184(outcome: Outcome4184): String = when (outcome) {
    is Outcome4184.Success -> outcome.payload
    Outcome4184.Failure -> "failed"
}
object Registry4185 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4185 = "registry-42"
    }
}
