package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16280<A, B>(val first: A, val second: B)

fun combine16280(left: Int, right: Int): Pair16280<Int, Int> {
    return Pair16280(left + 23, right + 59)
}
object Registry16281 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16281 = "registry-30"
    }
}
sealed class Outcome16282 {
    data class Success(val payload: String) : Outcome16282()
    data object Failure : Outcome16282()
}

fun describe16282(outcome: Outcome16282): String = when (outcome) {
    is Outcome16282.Success -> outcome.payload
    Outcome16282.Failure -> "failed"
}
data class Pair16283<A, B>(val first: A, val second: B)

fun combine16283(left: Int, right: Int): Pair16283<Int, Int> {
    return Pair16283(left + 44, right + 31)
}
object Registry16284 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16284 = "registry-51"
    }
}
