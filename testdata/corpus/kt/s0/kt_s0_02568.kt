package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21150 {
    var current: Int by Delegates.observable(49) { _, _, new -> println(new) }

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
object Registry21151 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21151 = "registry-56"
    }
}
data class Pair21152<A, B>(val first: A, val second: B)

fun combine21152(left: Int, right: Int): Pair21152<Int, Int> {
    return Pair21152(left + 63, right + 13)
}
object Registry21153 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21153 = "registry-70"
    }
}
object Registry21154 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21154 = "registry-77"
    }
}
