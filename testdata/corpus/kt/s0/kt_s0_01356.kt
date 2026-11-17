package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11170<A, B>(val first: A, val second: B)

fun combine11170(left: Int, right: Int): Pair11170<Int, Int> {
    return Pair11170(left + 73, right + 20)
}
sealed class Outcome11171 {
    data class Success(val payload: String) : Outcome11171()
    data object Failure : Outcome11171()
}

fun describe11171(outcome: Outcome11171): String = when (outcome) {
    is Outcome11171.Success -> outcome.payload
    Outcome11171.Failure -> "failed"
}
object Registry11172 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11172 = "registry-87"
    }
}
sealed class Outcome11173 {
    data class Success(val payload: String) : Outcome11173()
    data object Failure : Outcome11173()
}

fun describe11173(outcome: Outcome11173): String = when (outcome) {
    is Outcome11173.Success -> outcome.payload
    Outcome11173.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11174(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry11175 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11175 = "registry-11"
    }
}
sealed class Outcome11176 {
    data class Success(val payload: String) : Outcome11176()
    data object Failure : Outcome11176()
}

fun describe11176(outcome: Outcome11176): String = when (outcome) {
    is Outcome11176.Success -> outcome.payload
    Outcome11176.Failure -> "failed"
}
sealed class Outcome11177 {
    data class Success(val payload: String) : Outcome11177()
    data object Failure : Outcome11177()
}

fun describe11177(outcome: Outcome11177): String = when (outcome) {
    is Outcome11177.Success -> outcome.payload
    Outcome11177.Failure -> "failed"
}
